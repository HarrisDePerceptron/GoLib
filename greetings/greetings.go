package greetings

import
(

)
//This is a greeting stored in a vaiable
var Greeting = "Hello, how are you ";

//This is a greeting function
func GetGreeting(name string) string{
return Greeting + name;
}
