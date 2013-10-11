define ["devfeed", "header_view"], (Devfeed, HeaderView) ->

  Devfeed.module "HeaderApp.Common", (Common, Devfeed, Backbone, Marionette, $, _) ->

    headerView = new HeaderView.Header()

    Common.Controller =
      showHeader: ->
        Devfeed.headerRegion.show(headerView)

      showUserinfo: ->
        userSession = Devfeed.request("user:session")
        userinfoView = new HeaderView.UserinfoView
          model: userSession
        headerView.userinfoRegion.show(userinfoView)

      removeUserinfo: ->
        headerView.userinfoRegion.close()

  return Devfeed.HeaderApp.Common.Controller
