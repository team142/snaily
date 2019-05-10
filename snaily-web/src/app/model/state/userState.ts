
export class UserState {

    private static loggedIn: boolean = false;

    public static isLoggedIn(): boolean {
        return this.loggedIn;
    }

    public static login() {
        this.loggedIn = true;
    }

    public static logout() {
        this.loggedIn = false;
    }
}