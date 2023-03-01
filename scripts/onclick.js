function seeMore(div){
    var divContenu = div.nextSibling;
    if(divContenu.nodeType == 3) divContenu = divContenu.nextSibling;
    if(divContenu.style.display == 'flex') {
        divContenu.style.display = 'none';
    } else {
        divContenu.classList.add('cardAnimation')
        divContenu.style.display = 'flex';
    }
}

