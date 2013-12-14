define ["devfeed", "header_controller", "header_app"], (Devfeed, HeaderController, HeaderApp) ->

  userSessionHandler = null

  before ->
    # stub user:session request to return a fake authenticated user session
    userSessionHandler = Devfeed.reqres.getHandler "user:session"
    Devfeed.reqres.removeHandler "user:session"
    Devfeed.reqres.setHandler "user:session", ->
      userSession = new Devfeed.Entities.UserSession
        sessionId: 1234
        name: "Marconi Moreto"
      return userSession

  after ->
    # restore original user:session request handler
    Devfeed.reqres.removeHandler "user:session"
    Devfeed.reqres.setHandler "user:session", userSessionHandler

  describe "HeaderApp", ->
    it "shows header on Devfeed start", ->
      Backbone.history.constructor.started = false
      expect($('#header-region #header')[0]).to.not.exist
      Devfeed.start()
      expect($('#header-region #header')[0]).to.exist
      Devfeed.headerRegion.close()

  describe "HeaderApp.Common.Controller", ->
    it "showHeader displays header", ->
      expect($('#header-region #header')[0]).to.not.exist
      HeaderController.showHeader()
      expect($('#header-region #header')[0]).to.exist

    it "showUserinfo displays logged-in user info", ->
      expect($('#header-region #userinfo')[0]).to.not.exist
      HeaderController.showUserinfo()
      expect($('#header-region #userinfo')[0]).to.exist
      expect($("#user-name").html()).to.equal("Marconi Moreto")
