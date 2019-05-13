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

  public static setKey(v: string) {
    localStorage.setItem('key', v);
  }

  public static getKey(): string {
    return localStorage.getItem('key');
  }


}
