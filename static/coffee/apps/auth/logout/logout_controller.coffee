define ["devfeed"], (Devfeed) ->

  Devfeed.module "AuthApp.Logout", (Logout, Devfeed, Backbone, Marionette, $, _) ->

    Logout.Controller =
      logout: ->
        # clear projects data
        $("#projects-data").remove()
        window.Projects = null 

        Devfeed.request("session:logout")
        Devfeed.trigger("auth:login:show")

  return Devfeed.AuthApp.Logout.Controller
