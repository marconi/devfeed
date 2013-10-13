define [
  "devfeed",
  "login_controller",
  "register_controller",
  "activation_controller",
  "password_controller"
], (
  Devfeed,
  LoginController,
  RegisterController,
  ActivationController,
  PasswordController
) ->

  Devfeed.module "AuthApp", (AuthApp, Devfeed, Backbone, Marionette, $, _) ->

    class AuthApp.Router extends Marionette.AppRouter
      appRoutes:
        "login": "showLogin"
        "register": "showRegister"
        "activation/resend": "showActivation"
        "password/retrieve": "showForgotPassword"

    API =
      showLogin: ->
        LoginController.showLogin()
      showRegister: ->
        RegisterController.showRegister()
      showActivation: ->
        ActivationController.showActivation()
      showForgotPassword: ->
        PasswordController.showForgotPassword()

    Devfeed.on "auth:login:show", ->
      Devfeed.navigate("login")
      API.showLogin()

    Devfeed.on "auth:register:show", ->
      Devfeed.navigate("register")
      API.showRegister()

    Devfeed.on "auth:activation:show", ->
      Devfeed.navigate("activation/resend")
      API.showActivation()

    Devfeed.on "auth:forgotpass:show", ->
      Devfeed.navigate("password/retrieve")
      API.showForgotPassword()

    Devfeed.addInitializer ->
      new AuthApp.Router
        controller: API

  return Devfeed.AuthApp
