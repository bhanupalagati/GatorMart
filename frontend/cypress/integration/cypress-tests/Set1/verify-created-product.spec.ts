/// <reference types="cypress" />

//Test Objective 
//The create product page is accessed and the product is created.
//Then the details of the created products is verified.

describe('Start product creation', function () {
    it('Access the page', function() {
        cy.visit('http://localhost:4200/create')
        
    })

    it('Interacting with the 1st image field', function() {
        const main_image = 'cypress/fixtures/images/ipadimage.jpg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(0).find('input').selectFile(main_image)        
    })
    it('Interacting with the 2nd image field', function() {
        const detailed_images = 'cypress/fixtures/images/gadgets.jpeg'
        cy.get('div [class="ad_form_container__flex__input"]')
        .eq(1).find('input').selectFile(detailed_images)        
    })
    it('Text fields', function() {
        cy.get('input').eq(2)
        .type('Apple iPad M1') 
         cy.get('input').eq(3)
        .type('iPad Pro 12.9') 
        cy.get('input').eq(5)
        .clear().type('1200')
        cy.get('input').eq(6)
        .type('Simple Description for Cypress Iphone 1')
        cy.get('input').eq(7)
        .type('Description for Cypress Iphone 1')
        cy.get('input').eq(8)
        .type('Miami')
        cy.get('input').eq(9)
        .type('Florida')
        cy.get('input').eq(10)
        .type('12.1')
        cy.get('input').eq(11)
        .type('15.1')
        cy.get('input').eq(12)
        .clear().type('New')
        cy.get('input').eq(13)
        .clear().type('5')
        cy.get('input').eq(15)
        .type('Student')
        cy.get('input').eq(16)
        .type('Electronics')
    })
    it('Submit the form', function() {
        cy.get('button').click()       
    })

    it('Verify the details submitted by the user', function() {
        cy.get('mat-card').eq(2).click()
        cy.wait(1500) 
        cy.get('div [class="carousel-caption"]').find('h3').contains('Apple iPad M1')
        cy.wait(1500)
        cy.get('div [class="carousel-caption"]').find('p').contains('iPad Pro 12.9')
        cy.wait(1500)
        cy.get('div [class="details_container__main_info"]').find('div').eq(3).contains(1200)
        cy.wait(1500)
        cy.get('div [class="details_container__main_info"]').find('div').eq(4).contains('Miami, Florida')
        cy.wait(1500)
    })
})  