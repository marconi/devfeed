define ["devfeed"], (Devfeed) ->

  Devfeed.module "Common.Utils", (Utils, Devfeed, Backbone, Marionette, $, _) ->

    Utils.SmallSpin = lines: 8, length: 4, width: 3, radius: 5, corners: 1

  return Devfeed.Common.Utils
