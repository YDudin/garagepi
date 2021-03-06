"use strict";

$(document).ready(function(){

  var $btnLight = $("#btnLight");

  var lightOn = ($btnLight.text() == "Turn Off Light");

  function toggleGarageDoor() {
    $.post("/api/v1/toggle");
  }

  function turnLightOn() {
    $.post("/api/v1/light?state=on", parseLightState);
  }

  function turnLightOff() {
    $.post("/api/v1/light?state=off", parseLightState);
  }

  function parseLightState(data) {
    data = $.parseJSON(data);
    lightOn = data.LightOn;

    if (lightOn) {
      $btnLight.text("Turn Off Light");
    } else {
      $btnLight.text("Turn On Light");
    }

    if (!data.StateKnown) {
      $btnLight.prop('disabled', true);
    }
  }

  $("#btnDoorToggle").on("click", function() {
    toggleGarageDoor();
  });

  $btnLight.on("click", function() {
    if (lightOn) {
      turnLightOff();
    } else {
      turnLightOn();
    }
  });
});
