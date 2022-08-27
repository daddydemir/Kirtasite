package core

const (
	NoContent    string = "Aradığınız içerik bulunamadı."
	Ok           string = "İstek başarılı."
	Forbidden    string = "Yasaklı işlem."
	BadRequest   string = "Hatalı istek gönderdiniz."
	Created      string = "Model oluşturuldu."
	Updated      string = "Güncelleme başarılı."
	Unauthorized string = "Oturumunuzun süresi dolmuş."
	NotLogin     string = "İçeriğe ulaşabilmek için oturum açın."
	LoginFail    string = "Mail adresi yada Parola hatalı."
)

func SendMessage(message string) map[string]interface{} {
	return map[string]interface{}{"message": message}
}
