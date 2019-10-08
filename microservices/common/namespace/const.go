package namespace

import "fmt"

const Namespace = "knowledgemap"

func GetName(str string) string {
	return fmt.Sprintf("%s.%s", Namespace, str)
}
