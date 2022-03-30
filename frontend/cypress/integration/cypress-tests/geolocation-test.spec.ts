/// <reference types ="cypress" />

import { eq } from "cypress/types/lodash"

// Objective: Test the geolocation services

describe('Geolocation pull request test', function () {
    it('Page Load request', function () {
        cy.visit('http://localhost:4200/create')
    })

    it('Set first image', function() {
        cy.get('div [class="ad_form_container__flex__input"]').eq(0).find('input').selectFile('cypress/fixtures/images/headphones.jpg')        
    })
    it('Set second image', function() {
        cy.get('div [class="ad_form_container__flex__input"]').eq(1).find('input').selectFile('cypress/fixtures/images/ipadimage.jpeg')        
    })
    it('Test main geolocation button', function() {
        cy.get('input').eq(2).type('GeoTest Electricals') 
        cy.get('input').eq(3).type('GeoSync in Progress') 
        cy.get('input').eq(5).clear().type('999')
        cy.get('input').eq(6).type('Satellite downlink in progress')
        cy.get('input').eq(7).type('PSLV Satellite deployment')
        cy.get('input').eq(8).type('Marsville')
        cy.get('input').eq(9).type('Maston')
        cy.get('button').eq(0).click()
        cy.wait(5000)  
        //cy.get('input').eq(10).contains('29.62056457889113')
        //cy.get('input').eq(11).contains('-82.37617368280584'  x)
        cy.get('input').eq(12).clear().type('Used')
        cy.get('input').eq(13).clear().type('1')
        cy.get('input').eq(15).type('Scientists')
        cy.get('input').eq(16).type('Heavy Electronics')
    })
    it('Complete submission', function() {
        cy.get('button').click()})
})