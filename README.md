# statistics

pack 아카이브의 GEUL TAG 사용량을 Issuer별로 센다. TAG WORD `000K KKKC CCCC LLLL`에서 Length(LLLL)만 빼고 9비트 코드(`w>>4`, TAG ID `0x000L`~`0x1FFL`)로 센다. 그 밖에는 묶지 않는다: Reference의 양화, 자리 표지, 포인터도 각자의 코드다.

## 사용법

```bash
go run ./cmd/statistics [-memory 바이트] [-threads n] 아카이브...
```

- 아카이브는 `.well-known/geul/` 폴더이거나 pack set 폴더 하나(`.../0x00000000`)다. 여러 개를 주면 합친다. 같은 문서를 두 번 주면 두 번 센다.
- 코덱 문서(`0x00000000/codecs/`)가 든 `0x00000000` set이 인자 중 하나에 있어야 한다. 위키데이터 아카이브에는 없으므로 `0x00000000` set을 함께 준다.
- `-memory`는 프로세스 메모리 한도다(`debug.SetMemoryLimit`). 절반이 경로 추적 표의 상한이다. 0이면 한도가 없다.
- `-threads`는 스레드 수다(`GOMAXPROCS`). 기본값은 전체 코어다.

```bash
statistics -memory 1800000000 -threads 19 \
  draft/geul-org/artifacts/geul/.well-known/geul/0x00000000 \
  <위키데이터 아카이브>/.well-known/geul/0x00000001
```

출력은 표준 출력의 TSV 한 형식이다. 열은 `issuer`, `kind`, `codec`, `count`이고 첫 줄이 머리다. `kind`는 `REFERENCE`·`VALUE`·`EDGE`·`META`, `codec`은 TAG ID 표기 `0x000L`(9비트 코드 세 자리와 버린 Length 자리 L)이다.

```
issuer	kind	codec	count
0x00000000	EDGE	0x0A3L	9066
0x00000001	REFERENCE	0x014L	4759167
```

## 무엇을 세는가

- **문서.** 살아 있는 문서 주소(공표 주소 10가지 중 `hashes/`를 뺀 것: `<Issuer>/docs/<문서 ID 16자리>/`, `<Issuer>/codecs/<코덱 ID 8자리>/`, Issuer 목록 `issuers/`의 `<버전 ID>.geul`·`versions.geul`·`latest.geul`)마다 그 최신 레코드 하나. docs·codecs 목록은 없다(§1.1 101). 형식 밖 경로(옛 `<Issuer>/docs/<버전 ID>.geul` 같은 목록 경로 포함)는 문서로 세지 않는다. 버전 목록도 문서이고 버전마다 따로 센다. diff로 저장된 버전은 저장된 그대로 센다.
- **`<Issuer>/hashes/` 주소는 세지 않는다.** 내용 해시 조회일 뿐 문서의 주소가 아니고, 지난 버전도 그 주소로 남는다. hashes/ 전용 레코드는 건수만 표준 오류로 알린다.
- **레코드.** 문서의 레코드 전부. `DOC-*`, `BODY`, `END`도 TAG이므로 센다.
- **내장 레코드도 센다.** Edge의 구성요소(Edge Id 포함), Meta의 내장 Value·Reference, 값 안의 값(QUANTITY의 amount, UTF8-LANGUAGE-TEXT의 언어)은 레코드 안에 있지만 각자 TAG WORD가 있다. 자리 표지는 TRIPLE의 Object처럼 내장으로만 나온다. 경계는 `core.DecodeRecord`와 Codec 규격에 따른 `inspect.Layout`으로 찾고 깊이 제한 없이 내려간다.
- **규격이 없으면 열지 않는다.** 날값이 TAG처럼 보일 수 있어 Payload는 규격 없이 레코드로 읽지 않는다. 그 레코드 자신의 TAG만 세고, 열지 못한 건수를 표준 오류로 알린다. ISSUER 형식(각 Kind 구역의 끝 세 칸: Reference `0x05DL`~`0x05FL`, Value `0x09DL`~`0x09FL`, Edge `0x1BDL`~`0x1BFL`)은 아직 규격을 찾지 않으므로 여기에 든다.
- **Issuer**는 문서 주소의 첫 칸이다. 버전 목록처럼 DOC-ID가 없는 문서도 주소에는 Issuer가 있다. Issuer 목록 `issuers/`는 geul.org `0x00000000`의 것이다. 그 밖에 주소에 없으면 그 문서를 담은 set의 Issuer다.

## 경로 추적

pack의 덮어쓰기는 뒤 항목으로만 나타난다. 인벤토리에서 같은 경로가 뒤에 다시 나오면 앞 항목은 죽은 것이고, 삭제도 그런 뒤 항목이다. 경로 전체를 맵에 올리는 `Reader.State`나 `Archive.Paths`는 위키데이터(2.4억 경로)에서 쓸 수 없다.

1. 인벤토리를 한 번 읽어 pack별 단계 수와 문서 경로 수를 낸다.
2. 경로의 SHA-256 앞 128비트와 재생 순번을 한 줄(24바이트)로 하는 표를 만든다. 표가 `-memory`의 절반에 들어가도록 경로를 해시로 칸(k개)에 나누고, 칸마다 인벤토리 전체를 다시 읽어 표를 채우고 정렬해 같은 해시의 마지막이 아닌 줄을 죽은 항목으로 표시한다.
3. 남는 것은 인벤토리 단계당 1비트뿐이다(2.4억 단계에 30 MB). 레코드는 읽지 않는다.

128비트 해시가 부딪히면 그 문서를 빠뜨리지만 이 규모에서 확률은 2⁻⁷⁰ 아래다. 한 경로가 칸 크기를 넘게 되풀이되면 표가 차서 오류로 멈춘다.

## 메모리

pack을 차례로 하나씩 처리한다. 인벤토리는 통째로 메모리에 머무르므로(위키데이터 최대 380만 항목, 약 300 MB) 한 번에 하나만 연다. 작업자는 문서를 한 번에 하나씩 풀어 세고 표(`[512]uint64`)에 더한다.

gogeul의 인벤토리용 zstd 디코더는 처음 쓸 때의 `GOMAXPROCS` 수만큼 만들어지고 각각 마지막 인벤토리를 붙들고 있다. 그래서 프로세스는 스레드 1로 시작하고 문서를 셀 때 `-threads`로 올린다.

규격은 `Dir.CodecSpecs`로 주소로만 찾는다(`Dir.CodecSpec`은 `Archive.ReadDir`을 거쳐 모든 경로를 메모리에 올린다).

## 개발

`draft/go.work`에 묶여 있다. 답안 규칙은 filefunc(`filefunc validate .`)와 TSMA다. 라이선스는 CC0(`LICENSE`), 제3자 구성요소는 `NOTICE`.
