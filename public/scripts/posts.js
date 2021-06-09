const posts = document.body.getElementsByClassName("posts")


// see more
for (let post of posts) {
    const content = post.children[1].children[1];

    finalContent = "";
    for (let i = 0; i < 150; i++) {
        finalContent += content.textContent[i];
    }

    finalContent += " <a href='#' class='text-primary cursor-pointer' id='underline'>see more</a>"

    content.innerHTML = finalContent;
}



// like and dislikes
const likeSvgGreen = "<path d=\"M6.20081 0.335369C6.20048 0.335712 6.20008 0.335999 6.19974 0.336399L0.328771 6.31123C-0.111044 6.75884 -0.109407 7.48283 0.332609 7.92838C0.774569 8.37387 1.48942 8.37216 1.92929 7.92455L5.87097 3.91304L5.87097 28.1292C5.87097 28.7608 6.37644 29.2727 7 29.2727C7.62356 29.2727 8.12903 28.7608 8.12903 28.1292L8.12903 3.9131L12.0707 7.92449C12.5106 8.3721 13.2254 8.37382 13.6674 7.92832C14.1095 7.48277 14.111 6.75873 13.6712 6.31118L7.80026 0.33634C7.79992 0.335997 7.79952 0.335711 7.79919 0.335312C7.35785 -0.112528 6.64069 -0.111097 6.20081 0.335369Z\" fill=\"#34C931\"/>\n";
const likeSvgGrey = "<path d=\"M6.20081 0.335369C6.20048 0.335712 6.20008 0.335999 6.19974 0.336399L0.328771 6.31123C-0.111044 6.75884 -0.109407 7.48283 0.332609 7.92838C0.774569 8.37387 1.48942 8.37216 1.92929 7.92455L5.87097 3.91304L5.87097 28.1292C5.87097 28.7608 6.37644 29.2727 7 29.2727C7.62356 29.2727 8.12903 28.7608 8.12903 28.1292L8.12903 3.9131L12.0707 7.92449C12.5106 8.3721 13.2254 8.37382 13.6674 7.92832C14.1095 7.48277 14.111 6.75873 13.6712 6.31118L7.80026 0.33634C7.79992 0.335997 7.79952 0.335711 7.79919 0.335312C7.35785 -0.112528 6.64069 -0.111097 6.20081 0.335369Z\" fill=\"#C4C4C4\"/>\n";
const dislikeSvgRed = "<path d=\"M7.79919 28.9373C7.79952 28.937 7.79992 28.9367 7.80026 28.9363L13.6712 22.9615C14.111 22.5139 14.1094 21.7899 13.6674 21.3443C13.2254 20.8988 12.5106 20.9005 12.0707 21.3482L8.12903 25.3597L8.12903 1.14347C8.12903 0.51193 7.62356 0 7 0C6.37644 0 5.87097 0.51193 5.87097 1.14347L5.87097 25.3596L1.92929 21.3482C1.48942 20.9006 0.77457 20.8989 0.332609 21.3444C-0.109464 21.7899 -0.110988 22.514 0.328771 22.9615L6.19974 28.9364C6.20008 28.9367 6.20048 28.937 6.20081 28.9374C6.64215 29.3852 7.35931 29.3838 7.79919 28.9373Z\" fill=\"#FF1D1D\"/>\n";
const dislikeSvgGrey = "<path d=\"M7.79919 28.9373C7.79952 28.937 7.79992 28.9367 7.80026 28.9363L13.6712 22.9615C14.111 22.5139 14.1094 21.7899 13.6674 21.3443C13.2254 20.8988 12.5106 20.9005 12.0707 21.3482L8.12903 25.3597L8.12903 1.14347C8.12903 0.51193 7.62356 0 7 0C6.37644 0 5.87097 0.51193 5.87097 1.14347L5.87097 25.3596L1.92929 21.3482C1.48942 20.9006 0.77457 20.8989 0.332609 21.3444C-0.109464 21.7899 -0.110988 22.514 0.328771 22.9615L6.19974 28.9364C6.20008 28.9367 6.20048 28.937 6.20081 28.9374C6.64215 29.3852 7.35931 29.3838 7.79919 28.9373Z\" fill=\"#C4C4C4\"/>\n";

const likes = document.body.getElementsByClassName("likes");
for (let like of likes) {
    like.addEventListener("click", () => {
        like.firstElementChild.innerHTML = likeSvgGreen;
        like.children[1].classList.remove("txt-grey");
        like.children[1].classList.remove("txt-dark-2");
        like.children[1].classList.add("txt-green");

        const dislike = like.nextElementSibling;
        dislike.firstElementChild.innerHTML = dislikeSvgGrey;
        dislike.children[1].classList.remove("txt-red");
        dislike.children[1].classList.add("txt-grey");
        dislike.children[1].classList.add("txt-dark-2");
    })
}

const dislikes = document.body.getElementsByClassName("dislikes");
for (let dislike of dislikes) {
    dislike.addEventListener("click", () => {
        dislike.firstElementChild.innerHTML = dislikeSvgRed;
        dislike.children[1].classList.remove("txt-grey");
        dislike.children[1].classList.remove("txt-dark-2");
        dislike.children[1].classList.add("txt-red");

        const like = dislike.previousElementSibling;
        like.firstElementChild.innerHTML = likeSvgGrey;
        like.children[1].classList.remove("txt-red");
        like.children[1].classList.add("txt-grey");
        like.children[1].classList.add("txt-dark-2");
    })
}