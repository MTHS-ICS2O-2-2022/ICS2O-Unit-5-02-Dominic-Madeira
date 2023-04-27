// Copyright (c) 2023 Dominic M. All rights reserved
//
// Created by: Dominic M.
// Created on: April 2023
// This file contains the JS functions for index.html

"use strict"

/**
 * This function uses a selection component from https://github.com/CreativeIT/getmdl-select
 */

function myButtonClicked() {
  const positiveButton = document.getElementById("on-check").checked

  if (positiveButton == true) {
    const randomNumber = Math.floor(Math.random() * 6) + 1
    document.getElementById("answer").innerHTML =
      "Your number is: " + randomNumber
  } else {
    const randomNegativeNumber = Math.floor(Math.random() * 6 + 1) * -1
    document.getElementById("answer").innerHTML =
      "Your number is: " + randomNegativeNumber
  }
}
