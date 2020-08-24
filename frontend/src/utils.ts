var utils = {
  sanitizeHTML: function (text) {
    var element = document.createElement("div");
    element.innerText = text;
    return element.innerHTML;
  },

  humanizedBuildTime: function (time) {
    let human_time = moment(moment.unix(time)).format("L");
    return human_time;
  },

  humanizedRelativeBuildTime: function (time) {
    let human_time = moment(moment.unix(time)).fromNow();
    return human_time;
  },

  setItem(): function (tim)
};

module.exports = utils;
