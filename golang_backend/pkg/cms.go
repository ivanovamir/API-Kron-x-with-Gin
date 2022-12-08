package pkg

const (
	media_url = "https://api.kron-x.ru/media/"
)

func PhotoChecker(photo string) string {
	if photo == "" {
		return ""
	} else {
		return media_url + photo
	}
}
