import {likeAndDislike} from "./like-dislike.js";
import {minify} from "./posts-minify.js";

const postsContainer = document.getElementById("posts-container");

const likeSvgGrey = "<svg width=\"30\" height=\"30\" viewBox=\"0 0 14 30\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n" +
    "                                    <path d=\"M6.20081 0.335369C6.20048 0.335712 6.20008 0.335999 6.19974 0.336399L0.328771 6.31123C-0.111044 6.75884 -0.109407 7.48283 0.332609 7.92838C0.774569 8.37387 1.48942 8.37216 1.92929 7.92455L5.87097 3.91304L5.87097 28.1292C5.87097 28.7608 6.37644 29.2727 7 29.2727C7.62356 29.2727 8.12903 28.7608 8.12903 28.1292L8.12903 3.9131L12.0707 7.92449C12.5106 8.3721 13.2254 8.37382 13.6674 7.92832C14.1095 7.48277 14.111 6.75873 13.6712 6.31118L7.80026 0.33634C7.79992 0.335997 7.79952 0.335711 7.79919 0.335312C7.35785 -0.112528 6.64069 -0.111097 6.20081 0.335369Z\" fill=\"#C4C4C4\"/>\n" +
    "                                    </svg>"
const dislikeSvgGrey = "<svg width=\"30\" height=\"30\" viewBox=\"0 0 14 30\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n" +
    "<path d=\"M7.79919 28.9373C7.79952 28.937 7.79992 28.9367 7.80026 28.9363L13.6712 22.9615C14.111 22.5139 14.1094 21.7899 13.6674 21.3443C13.2254 20.8988 12.5106 20.9005 12.0707 21.3482L8.12903 25.3597L8.12903 1.14347C8.12903 0.51193 7.62356 0 7 0C6.37644 0 5.87097 0.51193 5.87097 1.14347L5.87097 25.3596L1.92929 21.3482C1.48942 20.9006 0.77457 20.8989 0.332609 21.3444C-0.109464 21.7899 -0.110988 22.514 0.328771 22.9615L6.19974 28.9364C6.20008 28.9367 6.20048 28.937 6.20081 28.9374C6.64215 29.3852 7.35931 29.3838 7.79919 28.9373Z\" fill=\"#C4C4C4\"/>\n" +
    "</svg>"
const moreSvg = "<svg width=\"30\" height=\"30\" viewBox=\"0 0 21 5\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"more-svg\"><circle cx=\"2.5\" cy=\"2.5\" r=\"2.5\" fill=\"#41599D\"/><circle cx=\"10.5\" cy=\"2.5\" r=\"2.5\" fill=\"#41599D\"/><circle cx=\"18.5\" cy=\"2.5\" r=\"2.5\" fill=\"#41599D\"/></svg"
const viewsSvg = "<svg width=\"30\" height=\"30\" viewBox=\"0 0 29 18\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n" +
    "                                <path d=\"M0.17944 8.4507C0.431738 8.10557 6.44314 0 14.1208 0C21.7984 0 27.8101 8.10557 28.0621 8.45037C28.3011 8.77779 28.3011 9.22188 28.0621 9.5493C27.8101 9.89443 21.7984 18 14.1208 18C6.44314 18 0.431738 9.89438 0.17944 9.54958C-0.0598392 9.22221 -0.0598392 8.77779 0.17944 8.4507ZM14.1208 16.1379C19.7762 16.1379 24.6744 10.7581 26.1243 8.99937C24.6762 7.23908 19.7883 1.86206 14.1208 1.86206C8.46565 1.86206 3.5678 7.24095 2.11723 9.00063C3.56532 10.7609 8.45324 16.1379 14.1208 16.1379Z\" fill=\"#C4C4C4\"/>\n" +
    "                                <path d=\"M14.1208 3.41379C17.201 3.41379 19.707 5.91983 19.707 9.00002C19.707 12.0802 17.201 14.5863 14.1208 14.5863C11.0406 14.5863 8.53456 12.0802 8.53456 9.00002C8.53456 5.91983 11.0406 3.41379 14.1208 3.41379ZM14.1208 12.7241C16.1744 12.7241 17.8449 11.0535 17.8449 9.00002C17.8449 6.9465 16.1743 5.2759 14.1208 5.2759C12.0673 5.2759 10.3967 6.9465 10.3967 9.00002C10.3967 11.0535 12.0672 12.7241 14.1208 12.7241Z\" fill=\"#C4C4C4\"/>\n" +
    "                                </svg>"
const commentsSvg = "<svg width=\"30\" height=\"30\" viewBox=\"0 0 21 19\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\">\n" +
    "                                <path d=\"M3.0567 18.4905H1.56871L2.6209 17.4383C3.18823 16.871 3.54251 16.1331 3.63776 15.3326C2.15732 14.3611 1.07476 13.0829 0.495268 11.6171C-0.0838141 10.1523 -0.156585 8.53992 0.284846 6.95409C0.814497 5.05125 2.05677 3.31687 3.78276 2.07041C5.65836 0.715953 7.99326 0 10.5351 0C13.7371 0 16.4361 0.919966 18.3401 2.66038C20.0553 4.22833 21 6.34328 21 8.61569C21 9.7197 20.7752 10.7939 20.3319 11.8085C19.8731 12.8584 19.2013 13.7932 18.3352 14.587C16.4286 16.3343 13.7314 17.2578 10.5351 17.2578C9.3485 17.2578 8.11033 17.0994 7.00488 16.8088C5.95852 17.8822 4.54063 18.4905 3.0567 18.4905ZM10.5351 1.2327C5.43832 1.2327 2.31092 4.27226 1.47239 7.28466C0.680009 10.1314 1.84808 12.8213 4.597 14.4801L4.90415 14.6655L4.8947 15.0241C4.87596 15.7339 4.70676 16.4175 4.40388 17.0355C5.14141 16.786 5.81064 16.3336 6.3347 15.7144L6.5953 15.4066L6.98175 15.5221C8.06624 15.8465 9.32816 16.0251 10.5351 16.0251C16.8783 16.0251 19.7673 12.1843 19.7673 8.61569C19.7673 6.69366 18.9651 4.9018 17.5083 3.57024C15.8354 2.04103 13.4241 1.2327 10.5351 1.2327Z\" fill=\"#C4C4C4\"/>\n" +
    "                                </svg>"

async function createPost(postContent) {
    const post = document.createElement("div");
    addClasses(post, [
        "card",
        "shadow-1",
        "rounded-2",
        "white",
        "posts",
        "mb-5",
    ])

    post.appendChild(await createHeader(postContent));
    post.appendChild(await createContent(postContent));
    post.appendChild(await createFooter(postContent));

    return post
}

function addClasses(element, classes) {
    for (const theClass of classes) {
        element.classList.add(theClass);
    }
}

async function createHeader(postContent) {
    const publisherId = postContent["PublisherID"];
    const url = "/user/" + publisherId

    const res = await fetch(url);
    const user = await res.json()

    const header = document.createElement("div");
    addClasses(header, [
        "card-header",
        "d-inline-flex",
        "vcenter",
        "w100",
        "pb-2",
    ])

    const img = document.createElement("img");
    img.src = user.ProfilePictureURL;
    img.alt = "";
    img.width = 40;
    img.height = 40;
    addClasses(img, [
        "rounded-4",
        "mr-4",
    ])
    header.appendChild(img);


    const name = document.createElement("p");
    addClasses(name, [
        "txt-grey",
        "txt-dark-3",
        "font-s2",
    ])
    const username = user["UserName"];
    name.textContent = "Posted by " + username
    header.appendChild(name);

    // Add clock system
    const time = document.createElement("p");
    addClasses(time, [
        "txt-grey",
        "txt-dark-3",
        "font-s2",
        "ml-auto",
    ])
    const clock = "1h";
    time.textContent = clock + " ago";
    header.appendChild(time);

    return header;
}

async function createContent(postContent) {
    const content = document.createElement("div");
    addClasses(content, [
        "card-content",
        "pt-0",
    ])

    // INSERT TITLE FROM API
    const title = document.createElement("h1");
    title.classList.add("font-s3");
    title.textContent = postContent.Title;
    content.appendChild(title);


    // INSERT CONTENT FROM API
    const text = document.createElement("p");
    addClasses(text, [
        "txt-grey",
        "txt-dark-5",
        "contents",
    ])
    text.textContent = postContent["Content"];
    content.appendChild(text);

    return content;
}

function createFooter(postContent) {
    const footer = document.createElement("div");
    addClasses(footer, [
        "card-footer",
        "pt-0",
        "vcenter",
        "d-flex",
        "fx-wrap",
    ])

    addLikesToFooter(footer);
    addDislikesToFooter(footer);
    addMoreToFooter(footer);

    const commentsAndViews = document.createElement("span");
    commentsAndViews.classList.add("ml-auto")
    addViewsToFooter(commentsAndViews);
    addCommentsToFooter(commentsAndViews);

    return footer;
}

function addLikesToFooter(footer) {
    const likes = document.createElement("div");
    addClasses(likes, [
        "likes",
        "d-flex",
        "vcenter",
        "mr-3",
        "cursor-pointer",
    ])
    likes.innerHTML = likeSvgGrey;

    // INSERT NUMBER OF LIKES FROM API
    const numberOfLikes = 1500;
    const likesContent = document.createElement("p");
    likesContent.classList.add("ml-2");
    likesContent.textContent = numberOfLikes.toString();

    likes.appendChild(likesContent);
    footer.appendChild(likes);
}

function addDislikesToFooter(footer) {
    const dislikes = document.createElement("div");
    addClasses(dislikes, [
        "dislikes",
        "d-flex",
        "vcenter",
        "mr-3",
        "cursor-pointer",
    ])
    dislikes.innerHTML = dislikeSvgGrey;

    // INSERT NUMBER OF LIKES FROM API
    const numberOfDislikes = 50;
    const dislikesContent = document.createElement("p");
    dislikesContent.classList.add("ml-2");
    dislikesContent.textContent = numberOfDislikes.toString();

    dislikes.appendChild(dislikesContent);
    footer.appendChild(dislikes);
}

function addMoreToFooter(footer) {
    const more = document.createElement("div");
    addClasses(more, [
        "d-inline-flex",
        "ml-5",
        "cursor-pointer",
    ])
    more.innerHTML = moreSvg;
    footer.appendChild(more);
}

function addViewsToFooter(footer) {
    const views = document.createElement("div");
    addClasses(views, [
        "d-inline-flex",
        "vcenter",
        "text-primary",
    ])
    views.innerHTML = viewsSvg;

    // INSERT VIEWS FROM API
    const viewsNumber = 3600;
    const viewsContainer = document.createElement("p");
    addClasses(viewsContainer, [
        "ml-1",
        "font-w600",
    ])
    viewsContainer.textContent = viewsNumber.toString();
    views.appendChild(viewsContainer);

    footer.appendChild(views);
}

function addCommentsToFooter(footer) {
    const comments = document.createElement("div");
    addClasses(comments, [
        "d-inline-flex",
        "vcenter",
        "text-primary",
        "ml-2"
    ])
    comments.innerHTML = commentsSvg;

    // INSERT COMMENTS FROM API
    const commentsNumber = 15;
    const commentsContainer = document.createElement("p");
    addClasses(commentsContainer, [
        "ml-1",
        "font-w600",
    ])
    commentsContainer.textContent = commentsNumber.toString();
    comments.appendChild(commentsContainer);

    footer.appendChild(comments);
}


async function main() {
    const res = await fetch("/posts")
    const postsList = await res.json();
    document.getElementById("nb-of-posts").textContent = postsList.length + " posts"
    for (let post of postsList) {
        const postToAdd = await createPost(post);
        postsContainer.appendChild(postToAdd);
    }
    likeAndDislike()
    minify();
}

main()
