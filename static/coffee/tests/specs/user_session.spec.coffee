define ["devfeed", "user_session"], (Devfeed, UserSession) ->

  describe "Entities.UserSession", ->
    it "returns false when no session", sinon.test ->
      @stub(UserSession.prototype, "isAuthenticated").returns(false)
      userSession = new UserSession()
      expect(userSession.isAuthenticated()).to.be.false

    it "returns true when there's session", sinon.test ->
      userSession = new UserSession
        sessionId: 1234
      expect(userSession.isAuthenticated()).to.be.true

    it "request session:authenticated returns false when not authenticated", sinon.test ->
      @stub(UserSession.prototype, "isAuthenticated").returns(false)
      expect(Devfeed.request("session:authenticated")).to.be.false

    it "request session:authenticated returns true when authenticated", sinon.test ->
      @stub(UserSession.prototype, "isAuthenticated").returns(true)
      expect(Devfeed.request("session:authenticated")).to.be.true

    it "request user:session returns UserSession instance", ->
      expect(Devfeed.request("user:session")).to.be.instanceof(UserSession)
