import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {ViewComponent} from './components/view/view.component';
import {LoginComponent} from './components/login/login.component';
import {WelcomeComponent} from './components/welcome/welcome.component';
import {HomeComponent} from './components/home/home.component';
import {HelpComponent} from './components/help/help.component';
import {NewComponent} from './components/new/new.component';


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
    path: 'view',
    component: ViewComponent
  },
  {
    path: 'home',
    component: HomeComponent
  },
  {
    path: 'new',
    component: NewComponent
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

export const routedComponents = [HelpComponent, LoginComponent, ViewComponent, WelcomeComponent, HomeComponent, NewComponent];
