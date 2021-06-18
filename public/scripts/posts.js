const contents = document.body.getElementsByClassName("contents")


// see more
for (let content of contents) {

    finalContent = "";
    for (let i = 0; i < 150; i++) {
        if (content.textContent.length > i) {
            finalContent += content.textContent[i];
        }

    }

    finalContent += " <span class='text-primary cursor-pointer post-details' id='underline'>see more</span>"

    content.innerHTML = finalContent;
}