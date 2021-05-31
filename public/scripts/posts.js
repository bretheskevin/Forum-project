const posts = document.body.getElementsByClassName("posts")

for (let post of posts) {
    const content = post.children[1].children[1];

    finalContent = "";
    for (let i = 0; i < 150; i++) {
        finalContent += content.textContent[i];
    }

    finalContent += " <a href='#' class='text-primary cursor-pointer' id='underline'>see more</a>"

    content.innerHTML = finalContent;
}