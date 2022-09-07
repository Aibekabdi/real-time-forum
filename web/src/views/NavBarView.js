import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
    }
    async getHtml(){
        return `
            <a href="/" class"nav__link" id="home-button" data-link>Home</a>
        `;
    }
}