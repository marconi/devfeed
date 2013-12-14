define ["devfeed", "register_view"], (Devfeed, RegisterView) ->

  Devfeed.module "AuthApp.Register", (Register, Devfeed, Backbone, Marionette, $, _) ->

    Register.Controller =
      showRegister: ->
        registerView = new RegisterView.RegisterForm()
        registerView.on "auth:register", (data) ->
          registering = Devfeed.request("session:register", data.name, data.email, data.password)
          $.when(registering).done (result) ->
            if result and result.errors?
              registerView.triggerMethod("form:data:invalid", result)
            else
              registerView.triggerMethod("form:data:valid", result)
        Devfeed.contentRegion.show(registerView)

  return Devfeed.AuthApp.Register.Controller
