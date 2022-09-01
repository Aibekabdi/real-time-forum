import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor() {
        super();
        this.setTitle("NewPost");
    }
    async getHtml(){
        // new post
        return `
        `;
    }

}