import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {BrowseComponent} from './components/browse/browse.component';
import {LoginComponent} from './components/login/login.component';
import {WelcomeComponent} from './components/welcome/welcome.component';
import {AccountComponent} from './components/account/account.component';
import {HelpComponent} from './components/help/help.component';


const routes: Routes = [
  {
    path: '',
    component: WelcomeComponent
  },
  {
    path: 'login',
    component: LoginComponent
  },
  {
    path: 'browse',
    component: BrowseComponent
  },
  {
    path: 'account',
    component: AccountComponent
  },
  {
    path: 'help',
    component: HelpComponent
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}

export const routedComponents = [HelpComponent, LoginComponent, BrowseComponent, WelcomeComponent, AccountComponent];
