define [
  "devfeed",
  "common_view",
  "tpl!apps/auth/register/templates/register.tpl",
  "syphon",
  "jquery_spin"
], (Devfeed, CommonView, registerTpl) ->

  Devfeed.module "AuthApp.Register.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.RegisterForm extends CommonView.FormViewMixin
      id: "register"
      className: "row collapse"
      template: registerTpl
      events:
        "click .login": "loginClicked"
        "click .signup": "signupClicked"
      ui:
        nameInput: "input[name=name]"
        emailInput: "input[name=email]"
        passwordInput: "input[name=password]"
        primaryBtn: ".signup"

      loginClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:login:show")

      signupClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("auth:register", data)

      clearFields: ->
        @ui.nameInput.val("")
        @ui.emailInput.val("")
        @ui.passwordInput.val("")
        @ui.nameInput.focus()

  return Devfeed.AuthApp.Register.View
