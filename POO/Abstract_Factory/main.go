package main

import "fmt"

// La mision es crear: SIS de notificaciones, tipo-> SMS y Email

// Interfaces global
type IFabricaNotificaciones interface {
	EnviarNotificacion()
	EnviarMensaje() IEnviar
}
// Interface de como enviar el mensaje
type IEnviar interface {
	MetodoEnviarMensage() string
	CanalDeEnvio() string
}
// Fin interfaces

// SMS
// Struct SMS notificacion 
type SMSNotification struct {

}

// Met Enviar Mensaje
func (SMSNotification) EnviarMensaje() IEnviar{
	// Se devuleve el STRUCT 
	return SMSEnviarNotificacion{}
}

// MetS EnviarNotificacion de SMS
func (SMSNotification) EnviarNotificacion() {
	fmt.Println("El mensaje se esta enviando por SMS")
}


// Struct Enviar notificaciones
type SMSEnviarNotificacion struct {

}
// Metodos -> SMSEnviarNotificacion
// Met enviar mensaje -> SMSEnviarNotificacion
func (SMSEnviarNotificacion) MetodoEnviarMensage() string {
	return "SMS"
}
// Met canal de envio -> SMSEnviarNotificacion
func (SMSEnviarNotificacion) CanalDeEnvio() string {
	return "TWilio"
}
// Fin metodos -> SMSEnviarNotificacion
//Fin SMS



// EMAIL
// Struct SMS notificacion 
type EmailNotification struct {

}

// Met Enviar Mensaje -> EmailNotification
func (EmailNotification) EnviarMensaje() IEnviar{
	// Se devuleve el STRUCT 
	return EmailEnviarNotificacion{}
}

// MetS EnviarNotificacion de SMS -> EmailNotification
func (EmailNotification) EnviarNotificacion() {
	fmt.Println("El mensaje se esta enviando por email")
}

// Struct Enviar notificaciones
type EmailEnviarNotificacion struct {

}

// Metodos -> EmailEnviarNotificacion
// Met enviar mensaje -> EmailEnviarNotificacion
func (EmailEnviarNotificacion) MetodoEnviarMensage() string {
	return "Email"
}
// Met canal de envio -> EmailEnviarNotificacion
func (EmailEnviarNotificacion) CanalDeEnvio() string {
	return "SES" // Servicio de Amazon
}
// Fin metodos -> EmailEnviarNotificacion
// Fin EMAIL



// Func para conocer el tipo de notificacion
func ObtenerNotificacion (notificacion string) (IFabricaNotificaciones, error) {
	if notificacion == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificacion == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No se genero la notificación")
}

// func para realizar el envio de la notificiacion 
func enviarNotificacion(f IFabricaNotificaciones){
	f.EnviarNotificacion()
}

// func para obtener el metodo usado
func obtenerMetodo(f IFabricaNotificaciones){
	fmt.Println(f.EnviarMensaje().MetodoEnviarMensage()) 
}


func main () {

	smsFactory, _ := ObtenerNotificacion("SMS")
	emailFactory, _ := ObtenerNotificacion("Email")

	enviarNotificacion(smsFactory)
	enviarNotificacion(emailFactory)

	obtenerMetodo(smsFactory)
	obtenerMetodo(emailFactory)

}