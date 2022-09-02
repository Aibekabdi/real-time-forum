import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("NewPost");
    }
    async getHtml(){
        // new post
        return `
        `;
    }

}