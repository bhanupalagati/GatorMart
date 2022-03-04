/// <reference types="cypress" />

//Testing all the applicaiton functionality and 
//adding multiple products and checking the 
//output for the flaws

describe('Access main page', function () {
    it('Hitting the backend', function() {
        cy.visit('http://localhost:4200/')
        cy.wait(2000)      
    })
    //Creating 3 products for testing
    for (let i = 0; i < 3; i++) 
    { 
        it('Access the create new product functionality', function() {
          cy.visit('http://localhost:4200/create')
          cy.wait(2000)
         })
        it('Set the main image', function() {
          const image_single = 'cypress/fixtures/images/ipadimage.jpg'
          cy.get('div [class="ad_form_container__flex__input"]').eq(0)
          .find('input').selectFile(image_single)     
          cy.wait(1000)   
        })
        it('Set the carousel images', function() {
           cy.get('div [class="ad_form_container__flex__input"]').eq(1)
           .find('input').selectFile(
               [   'cypress/fixtures/images/headphones.jpg',
                    'cypress/fixtures/images/gadgets.jpeg' ])   
            cy.wait(1000)      
        })
        it('Fill the form', function() {
          cy.get('input').eq(2).type('Electronics') 
          cy.wait(1000) 
          cy.get('input').eq(3).type('Apple Electronics') 
          cy.wait(1000) 
          cy.get('input').eq(5).clear().type('2500')
          cy.wait(1000) 
          cy.get('input').eq(6).type('Simple Description for Apple electronics')
          cy.wait(1000) 
          cy.get('input').eq(7).type('Description for Cypress Iphone 1')
          cy.wait(1000) 
          cy.get('input').eq(8).type('Jacksonville')
          cy.wait(1000) 
          cy.get('input').eq(9).type('Florida')
          cy.wait(1000) 
          cy.get('input').eq(10).type('78.1')
          cy.wait(1000) 
          cy.get('input').eq(11).type('45.1')
          cy.wait(1000) 
          cy.get('input').eq(12).clear().type('New')
          cy.wait(1000) 
          cy.get('input').eq(13).clear().type('5')
          cy.wait(1000) 
          cy.get('input').eq(15).type('Student')
          cy.wait(1000) 
          cy.get('input').eq(16).type('Electronics')
          cy.wait(1000) 
        })
        it('List the product on the website', function() {
          cy.get('button').click()       
        })
        it('Multi-product wait time', function() {
          cy.wait(2500)       
        })
    }
    it('Verify the details submitted by the user', function() {
        cy.get('mat-card').eq(3).click()
        // This is for the heading
        cy.get('div [class="carousel-caption"]').find('h3').contains('Macbook Pro')
        //This is for the secondary heading
        cy.get('div [class="carousel-caption"]').find('p').contains('Apple Macbook')
        // This is for the price
        cy.get('div [class="details_container__main_info"]').find('div').eq(3).contains(2400)
        // This is for the city
        cy.get('div [class="details_container__main_info"]').find('div').eq(4).contains('Clearwater, Florida')
    })
    it('Remove the product from the store', function() {
        cy.get('mat-card').eq(0)      
        .click()  
        cy.wait(1000)  
        cy.scrollTo('bottom')
        cy.wait(1000)
        cy.get('button').eq(1)
        .click()
        cy.wait(1000)
        cy.get('button').eq(3)
        .click()
        cy.wait(1000)
    })

})

