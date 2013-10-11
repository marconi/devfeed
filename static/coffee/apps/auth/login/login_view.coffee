define [
  "devfeed",
  "common_view",
  "tpl!apps/auth/login/templates/login.tpl",
  "syphon",
  "jquery_spin"
], (Devfeed, CommonView, loginTpl) ->

  Devfeed.module "AuthApp.Login.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.LoginForm extends CommonView.FormViewMixin
      id: "login"
      className: "row collapse"
      template: loginTpl
      events:
        "click .signin": "signinClicked"
        "click .register": "registerClicked"
        "click .forgot-pass": "forgotPassClicked"
        "click .resend-activation": "resendActivationClicked"
      ui:
        primaryBtn: ".signin"

      signinClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("auth:login", data)

      registerClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:register:show")

      forgotPassClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:forgotpass:show")

      resendActivationClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:activation:show")

  return Devfeed.AuthApp.Login.View
