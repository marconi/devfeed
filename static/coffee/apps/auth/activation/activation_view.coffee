define [
  "devfeed",
  "common_view",
  "tpl!apps/auth/activation/templates/resend.tpl",
  "syphon",
  "jquery_spin"
], (Devfeed, CommonView, resendTpl) ->

  Devfeed.module "AuthApp.Activation.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.ResendForm extends CommonView.FormViewMixin
      id: "activation"
      className: "row collapse"
      template: resendTpl
      events:
        "click .send": "sendClicked"
        "click .login": "loginClicked"
      ui:
        emailInput: "input[name=email]"
        primaryBtn: ".send"

      sendClicked: (e) ->
        e.preventDefault()
        if @ui.primaryBtn.hasClass("disabled")
          return

        @clearErrors()
        @showPreloader()

        data = Backbone.Syphon.serialize(@)
        @trigger("auth:activation", data)

      loginClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("auth:login:show")

      clearFields: ->
        @ui.emailInput.val("").focus()

  return Devfeed.AuthApp.Activation.View
