import AbstractView from "./AbstractView.js";

export default class extends AbstractView{
    constructor(params) {
        super(params);
        this.setTitle("Signup");
    }
    async getHtml(){
        //signup
        return `
        <div class="signup-container">
            <h3> Sign up </h3>
            First name: <br>
            <input type="text" id="firstname" placeholder="type your name" required>
            Last name: <br>
            <input type="text" id="lastname" placeholder="type your surname" required>
            Email: <br>
            <input type="text" id="email" placeholder="type your email" required>
            Username: <br>
            <input type="text" id="nickname" placeholder="type your username" required>
            Age: <br>
            <input type="number" id="age" min="10" max="99" placeholder="type your age" required>
            Gender: <br>
            <input type="radio" name="gender" id="gender-male" value="1" required> Male
            <input type="radio" name="gender" id="gender-female" value="2"> Female <br> <br>
            Password: <br>
            <input type="password" id="password" placeholder="type your password"  minlength="7" maxlength="64" required>
            Confirm password: <br>
            <input type="password" id="password-confirm" placeholder="Password" maxlength="64" required>
            <div class="error" id="error-message"></div>
            <button type="submit">Sign up</button>
        </div>
        `;
    }

}