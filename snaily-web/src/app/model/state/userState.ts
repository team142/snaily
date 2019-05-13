export class UserState {

  public static isLoggedIn(): boolean {
    return localStorage.getItem('login') === 'true';
  }

  public static login() {
    localStorage.setItem('login', 'true');
  }

  public static logout() {
    localStorage.setItem('login', 'false');
  }


}
