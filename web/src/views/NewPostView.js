import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("Create Post");
    }
    async getHtml(){
        // new post
        return `
                <p>Title:</p>
                <input type="text" id="title-input" minlength="2" maxlength="64" required>
                <p>Categories:</p>
                <select id="categories-select"></select>
                <button id='createPost'>create</button>
        `;
    }
    
}