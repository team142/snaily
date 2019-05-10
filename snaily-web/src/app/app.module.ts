import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';
import {AppComponent} from './app.component';
import {FormsModule} from '@angular/forms';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';

// Libs
import {HttpModule} from '@angular/http';

// Routing
import {AppRoutingModule, routedComponents} from './app-routing.module';
import {WelcomeComponent} from './components/welcome/welcome.component';
import {HomeComponent} from './components/home/home.component';
import {NewComponent} from './components/new/new.component';
import {HelpComponent} from './components/help/help.component';
import {ApiComponent} from './components/api/api.component';
import {ApiIssuesComponent} from './components/api/api-issues/api-issues.component';
import {ApiDocsComponent} from './components/api/api-docs/api-docs.component';
import {ApiAccessComponent} from './components/api/api-access/api-access.component';
import {ApiVersionsComponent} from './components/api/api-versions/api-versions.component';
import {ViewComponent} from './components/view/view.component';

// import './rxjs-extensions';

@NgModule({
  declarations: [
    AppComponent,
    routedComponents,
    WelcomeComponent,
    HomeComponent,
    NewComponent,
    ViewComponent,
    HelpComponent,
    ApiComponent,
    ApiIssuesComponent,
    ApiDocsComponent,
    ApiAccessComponent,
    ApiVersionsComponent
  ],
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
