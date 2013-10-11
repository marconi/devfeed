define [
  "devfeed",
  "common_view",
  "tpl!apps/auth/password/templates/retrieve.tpl",
  "syphon",
  "jquery_spin"
], (Devfeed, CommonView, retrieveTpl) ->

  Devfeed.module "AuthApp.Password.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.ForgotPasswordForm extends CommonView.FormViewMixin
      id: "forgot-password"
      className: "row collapse"
      template: retrieveTpl
      events:
        "click .retrieve": "retrieveClicked"
        "click .login": "loginClicked"
      ui:
        emailInput: "input[name=email]"
        primaryBtn: ".retrieve"

      retrieveClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("auth:forgotpass", data)

      loginClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:login:show")

      clearFields: ->
        @ui.emailInput.val("").focus()

  return Devfeed.AuthApp.Password.View
