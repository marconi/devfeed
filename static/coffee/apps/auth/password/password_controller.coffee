define ["devfeed", "password_view"], (Devfeed, PasswordView) ->

  Devfeed.module "AuthApp.Password", (Password, Devfeed, Backbone, Marionette, $, _) ->

    Password.Controller =
      showForgotPassword: ->
        forgotPassView = new PasswordView.ForgotPasswordForm()
        forgotPassView.on "auth:forgotpass", (data) ->
          retrieving = Devfeed.request("session:password:retrieve", data.email)
          $.when(retrieving).done (result) ->
            if result and result.errors?
              forgotPassView.triggerMethod("form:data:invalid", result)
            else
              forgotPassView.triggerMethod("form:data:valid", result)
        Devfeed.contentRegion.show(forgotPassView)

  return Devfeed.AuthApp.Password.Controller
