package file

import "os"

func FileRead(file string) string {
 	conteudo, err := os.ReadFile(file)
    if err != nil {
        return err.Error()
    }
		return string(conteudo)
}
