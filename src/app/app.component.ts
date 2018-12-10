import { Component,OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http'; 

const BASE_URL = 'https://play-zik.appspot.com';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  isAdminMode: boolean = true;

  constructor(private _http: HttpClient) {}

  ngOnInit() {
    
  }

}
