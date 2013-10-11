define ["devfeed", "activation_view"], (Devfeed, ActivationView) ->

  Devfeed.module "AuthApp.Activation", (Activation, Devfeed, Backbone, Marionette, $, _) ->

    Activation.Controller =
      showActivation: ->
        resendView = new ActivationView.ResendForm()
        resendView.on "auth:activation", (data) ->
          resending = Devfeed.request("session:activation:resend", data.email)
          $.when(resending).done (result) ->
            if result and result.errors?
              resendView.triggerMethod("form:data:invalid", result)
            else
              resendView.triggerMethod("form:data:valid", result)
        Devfeed.contentRegion.show(resendView)

  return Devfeed.AuthApp.Activation.Controller
