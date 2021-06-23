export function minify() {
    const contents = document.body.getElementsByClassName("contents")

    // see more
    for (let content of contents) {

        let finalContent = "";
        for (let i = 0; i < 150; i++) {
            finalContent += content.textContent[i];
        }

        finalContent += " <a href='#' class='text-primary cursor-pointer' id='underline'>see more</a>"

        content.innerHTML = finalContent;
    }
}

