"use strict"

let runButton = document.querySelector(".run");
// got the run button element

let inputArea = document.querySelector("#input");
// got the input element for the input area

let outputArea = document.querySelector("#output");
// got the input element for the output area

runButton.addEventListener( "click", function() {
  outputArea.value = inputArea.value;
} )
// Added a click event to the run button so that whenever it is clicked whatever value( this is an input element property) in the input will be the value of the output 
