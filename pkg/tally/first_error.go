//ff:func feature=count type=helper control=sequence
//ff:what firstError — 작업자의 오류를 앞세워 둘 중 하나를 고른다
package tally

// firstError picks the cause over what it left behind in the inventory walk.
func firstError(cause, err error) error {
	if cause != nil {
		return cause
	}
	return err
}
