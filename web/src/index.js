import Post from "./views/PostView.js";
import Chats from "./views/ChatsView.js";
import Signin from "./views/SigninView.js";
import Signup from "./views/SignupView.js";
import Home from "./views/HomeView.js";
import Profile from "./views/ProfileView.js";
import NewPost from "./views/NewPostView.js";
import NavBar from "./views/NavBarView.js";

const pathToRegex = path => new RegExp("^" + path.replace(/\//g, "\\/").replace(/:\w+/g, "(.+)") + "$");

const getParams = match => {
    const values = match.result.slice(1);
    const keys = Array.from(match.route.path.matchAll(/:(\w+)/g)).map(result => result[1]);
    
    return Object.fromEntries(keys.map((key, i) => {
        return [key, values[i]];
    }));
};

const navigateTo = url => {
    history.pushState(null , null, url);
    router();
};

const router = async () => {
    const routes = [
        {path: "/", view: Home},
        {path: "/sign-in", view: Signin},
        {path: "/sign-up", view: Signup},
        {path: "/chats", view: Chats},
        {path: "/post/:postID", view: Post},
        {path: "/user/:userID", view: Profile},
        {path: "/create-post", view: NewPost},
    ];

    // Test each route for potential match
    const potentialMatches = routes.map(route => {
        return {
            route : route,
            result: location.pathname.match(pathToRegex(route.path))
        };
    });
    
    let match = potentialMatches.find(potentialMatch => potentialMatch.result !== null)
    
    if (!match) {
        match = {
            route: routes[0],
            result: [location.pathname]
        };
    }
    
    const NavBarView = new NavBar(null);
    document.querySelector("#navbar").innerHTML = await NavBarView.getHtml();

    const view = new match.route.view(getParams(match));

    document.querySelector("#app").innerHTML = await view.getHtml();
    view.init();
};

window.addEventListener("popstate", router);

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")){
            e.preventDefault();
            navigateTo(e.target.href);
        }
    })
    
    router();
})