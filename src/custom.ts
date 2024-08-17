htmx.on("htmx:beforeSwap", function (evt) {
  if (evt.detail.xhr.status === 404) {
    evt.detail.shouldSwap = true;
    evt.detail.isError = false;
  }
});
