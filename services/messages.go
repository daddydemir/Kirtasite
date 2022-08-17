package services

const (
	NoContent string = "Aradığınız içerik bulunamadı."
	Ok        string = "İstek başarılı."
)

func SendMessage(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
