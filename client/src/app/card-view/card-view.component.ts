import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';

import { Router } from '@angular/router';


@Component({ templateUrl: 'card-view.component.html',
styleUrls: ["./card-view.component.css"] })

export class CardViewComponent {

    position = 0;
    frontItems: string[] = [''];
    backItems: string[] = [''];
    cards: string[][] = [[''], ['']];
    size = this.frontItems.length;
    showFront: boolean = true;

    constructor(private http: HttpClient, private router:Router) { }

    ngOnInit(): void {

        const folderId = localStorage.getItem("folderId");

        this.http.get('http://api.memorly.kro.kr/folders/' + folderId, { "headers": {"Authorization": sessionStorage.getItem('accessToken') || ""} }).subscribe((response: any) => {

        console.log(response);
          this.cards = response.data.cards.map((card: any) => ({
            
            //id: folder.id,
            //title: folder.title,
            //icon: folder.type === 'document' ? 'description' : 'folder'
          }));
        });
      }

    cycleLeft() {
        if (this.position != 0) {
            this.position--;
        }
        else {
            this.position = this.size - 1;
        }
      }
    
      cycleRight() {
        if (this.position != this.size - 1) {
            this.position++;
        }
        else {
            this.position = 0;
        }
      }

      flipCard() {
        if (this.showFront) {
            this.showFront = false;
        }
        else {
            this.showFront = true;
        }
      }

}