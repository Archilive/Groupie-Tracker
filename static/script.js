window.onload = function(){
    slideralbmin();
    slideralbmax();
    slidercreamin();
    slidercreamax();
}
// RANGE 1 / DATE FIRST ALBUM
let sliderOneMin = document.getElementById("slider_alb_min");
let sliderOneMax = document.getElementById("slider_alb_max");
let displayValOneMin = document.getElementById("range_alb_min");
let displayValOneMax = document.getElementById("range_alb_max");
let minGap = 0;
let sliderTrackOne = document.querySelector(".slider-trackOne");
let sliderOneMaxValue = (((document.getElementById("slider_alb_min").max-1963)/55)*100);


function slideralbmin(){
    if(parseInt(sliderOneMax.value) - parseInt(sliderOneMin.value) <= minGap){
        sliderOneMin.value = parseInt(sliderOneMax.value) - minGap;
    }
    displayValOneMin.textContent = sliderOneMin.value;
    fillColorOne();
}
function slideralbmax(){
    if(parseInt(sliderOneMax.value) - parseInt(sliderOneMin.value) <= minGap){
        sliderOneMax.value = parseInt(sliderOneMin.value) + minGap;
    }
    displayValOneMax.textContent = sliderOneMax.value;
    fillColorOne();
}


function fillColorOne(){
    percentOne1 = (((sliderOneMin.value-1963)) / 55) * 100;
    percentOne2 = (((sliderOneMax.value-1963)) / 55) * 100;
    sliderTrackOne.style.background = `linear-gradient(to right, #dadae5 ${percentOne1}% , #3264fe ${percentOne1}% , #3264fe ${percentOne2}%, #dadae5 ${percentOne2}%)`;
}

// RANGE 2 / DATE CREATION
let sliderTwoMin = document.getElementById("slider_crea_min");
let sliderTwoMax = document.getElementById("slider_crea_max");
let displayValTwoMin = document.getElementById("range_crea_min");
let displayValTwoMax = document.getElementById("range_crea_max");
let sliderTrackTwo = document.querySelector(".slider-trackTwo");
let sliderTwoMaxValue = (((document.getElementById("slider_crea_min").max-1958)/57)*100);

function slidercreamin(){
    if(parseInt(sliderTwoMax.value) - parseInt(sliderTwoMin.value) <= minGap){
        sliderTwoMin.value = parseInt(sliderTwoMax.value) - minGap;
    }
    displayValTwoMin.textContent = sliderTwoMin.value;
    fillColorTwo();
}
function slidercreamax(){
    if(parseInt(sliderTwoMax.value) - parseInt(sliderTwoMin.value) <= minGap){
        sliderTwoMax.value = parseInt(sliderTwoMin.value) + minGap;
    }
    displayValTwoMax.textContent = sliderTwoMax.value;
    fillColorTwo();
}

function fillColorTwo(){
    percentTwo1 = (((sliderTwoMin.value-1958)) / 57) * 100;
    percentTwo2 = (((sliderTwoMax.value-1958)) / 57) * 100;
    sliderTrackTwo.style.background = `linear-gradient(to right, #dadae5 ${percentTwo1}% , #3264fe ${percentTwo1}% , #3264fe ${percentTwo2}%, #dadae5 ${percentTwo2}%)`;
}

const targetDivCity = document.getElementById("dropdown-child");
const btnCity = document.getElementById("city");
targetDivCity.setAttribute("style","display:none;");
btnCity.onclick = function () {
    if (targetDivCity.style.display == "none") {
    targetDivCity.setAttribute("style","display:block;");
    } else {
    targetDivCity.setAttribute("style","display:none;");
    }
};

const targetDivOrder = document.getElementById("dropdown-child-order");
const btnOrder = document.getElementById("order");
targetDivOrder.setAttribute("style","display:none;");
btnOrder.onclick = function () {
    if (targetDivOrder.style.display == "none") {
    targetDivOrder.setAttribute("style","display:block;");
    } else {
    targetDivOrder.setAttribute("style","display:none;");
    }
};

// search bar
function searching() {
    let input = document.getElementById('searchbar').value
    input=input.toLowerCase();
    let x = document.getElementsByClassName('card');
    let y = document.getElementsByClassName('membersa');
    for (i = 0; i < x.length; i++) { 
        console.log(input)
        if (!x[i].innerHTML.toLowerCase().includes(input) && !y[i].innerHTML.toLowerCase().includes(input)) {
            x[i].style.display="none";
        }
        else {
            x[i].style.display="flex";       
        }
    }
}