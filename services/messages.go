package services

const (
	NoContent  string = "Aradığınız içerik bulunamadı."
	Ok         string = "İstek başarılı."
	Forbidden  string = "Yasaklı işlem."
	BadRequest string = "Hatalı istek gönderdiniz."
	Created    string = "Model oluşturuldu."
	Updated    string = "Güncelleme başarılı."
)

func SendMessage(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
