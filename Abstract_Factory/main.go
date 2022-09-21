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

// Struct SMS notificacion 
type SMSNotification struct {

}

func (SMSNotification) EnviarMensaje() IEnviar{
	// Se devuleve el STRUCT 
	return SMSEnviarNotificacion{}
}

// MetS EnviarNotificacion de SMS
func (SMSNotification) EnviarNotificacion() {
	fmt.Println("El mensaje se esta enviando")
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



func main () {

}