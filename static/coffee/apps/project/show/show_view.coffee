define [
  "devfeed",
  "jquery_livefilter",
  "common_utils",
  "tpl!apps/project/show/templates/sidebar.tpl",
  "tpl!apps/project/show/templates/findstory.tpl",
  "tpl!apps/project/show/templates/filterpreloader.tpl",
  "tpl!apps/project/show/templates/stories.tpl",
  "tpl!apps/project/show/templates/story.tpl",
  "tpl!apps/project/show/templates/empty.tpl",
  "tpl!apps/project/show/templates/chatinfo.tpl",
  "tpl!apps/project/show/templates/chatbox.tpl",
  "tpl!apps/project/show/templates/show.tpl"
], (
  Devfeed,
  LiveFilter,
  CommonUtils,
  sidebarTpl,
  findStoryTpl,
  filterPreloaderTpl,
  storiesTpl,
  storyTpl,
  emptyTpl,
  chatinfoTpl,
  chatboxTpl,
  showTpl
  ) ->

  Devfeed.module "ProjectApp.Show.View", (View, Devfeed, Backbone, Marionette, $, _) ->

    class View.Empty extends Marionette.ItemView
      className: "empty"
      tagName: "li"
      template: emptyTpl

    class View.Story extends Marionette.ItemView
      tagName: "li"
      template: storyTpl
      events:
        "click .name": "nameClicked"
        "click .task": "taskClicked"

      onRender: ->
        @$el.addClass(@model.get("current_state"))

      nameClicked: (e) ->
        e.preventDefault()
        @$el.toggleClass("open")
        @$(".tasks").toggleClass("hide")

      taskClicked: (e) ->
        e.preventDefault()
        $a = $(e.currentTarget)
        $a.toggleClass("complete")

        $checkbox = $a.find("input[type=checkbox]")
        $checkbox.prop "checked", (i, value) ->
          return not value

    class View.Stories extends Marionette.CompositeView
      id: "stories"
      template: storiesTpl
      emptyView: View.Empty
      itemView: View.Story
      itemViewContainer: ".inner ul"
      events:
        "click .more": "moreClicked"

      moreClicked: (e) ->
        e.preventDefault()
        @$(".more span").addClass("hide")
        @$(".more div").spin(CommonUtils.SmallSpin)
        @trigger("stories:more")

      onMoreStories: ->
        @$(".more div").spin(false)
        @$(".more span").removeClass("hide")
        @trigger("stories:more:rendered")

      onCompositeCollectionRendered: ->
        if @collection.size() > 0
          @$(".more").removeClass("hide")
        else
          @$(".more").addClass("hide")

    class View.FilterPreloader extends Marionette.ItemView
      id: "filter-preloader"
      template: filterPreloaderTpl
      onDomRefresh: ->
        @$(".loading").spin(CommonUtils.SmallSpin)
        if not $("#find-story .settings").hasClass("hide")
          @$el.addClass("settings-shown")

    class View.FindStory extends Marionette.ItemView
      id: "find-story"
      className: "small-12 columns"
      template: findStoryTpl
      events:
        "click .settings-cog a": "settingsClicked"
        "change .settings input[type=checkbox]": "settingsChanged"
      filters: []

      onRender: ->
        @bindLiveFilter()

      bindLiveFilter: ->
        @$('.keyword input').fastLiveFilter('#stories .inner ul:first-child')        

      settingsClicked: (e) ->
        e.preventDefault()
        if @$(".settings").hasClass("hide")
          @trigger("settings:shown")
        else
          @trigger("settings:hidden")  
        @$(".settings").toggleClass("hide")

      settingsChanged: (e) ->
        filters = _.filter @$(".settings input[type=checkbox]"), (checkbox) ->
          return $(checkbox).is(":checked")
        filters = _.map filters, (filter) ->
          return $(filter).attr("name")
        @filters = filters
        @trigger("filters:changed", @filters)

    class View.Sidebar extends Marionette.Layout
      id: "sidebar"
      template: sidebarTpl
      events:
        "click #hide-sidebar": "hidesidebarClicked"
      regions:
        findStoryRegion: "#find-story-region"
        storiesRegion: "#stories-region"
      filterPreloaderView: null

      hidesidebarClicked: (e) ->
        e.preventDefault()
        console.log "hiding..."

      onSettingsShown: ->
        @storiesRegion.currentView.$el.addClass("settings-shown")
        if @filterPreloaderView
          @filterPreloaderView.$el.addClass("settings-shown")        

      onSettingsHidden: ->
        @storiesRegion.currentView.$el.removeClass("settings-shown")
        if @filterPreloaderView
          @filterPreloaderView.$el.removeClass("settings-shown")

    class View.Chatinfo extends Marionette.ItemView
      id: "chatinfo"
      template: chatinfoTpl
      events:
        "click #back-projects a": "backprojectsClicked"

      backprojectsClicked: (e) ->
        e.preventDefault()
        Devfeed.trigger("projects:list")

    class View.Chatbox extends Marionette.ItemView
      id: "chatbox"
      template: chatboxTpl

    class View.Show extends Marionette.Layout
      id: "project-details"
      className: "row collapse"
      template: showTpl
      regions:
        sidebarRegion: "#sidebar-region"
        chatinfoRegion: "#chatinfo-region"
        chatboxRegion: "#chatbox-region"

  return Devfeed.ProjectApp.Show.View
