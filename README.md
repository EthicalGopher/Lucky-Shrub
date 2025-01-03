
---

# 🌿 Lucky Shrub - Garden Design Website  

Lucky Shrub is a modern, responsive website designed for a garden design and landscaping company based in Tucson, Arizona. The site showcases the company's services, plant nursery, and contact information with a sleek and user-friendly design.  

---

## 📋 Table of Contents  
- [Features](#features)  
- [Technologies Used](#technologies-used)  
- [Setup Instructions](#setup-instructions)  
- [Contributing](#contributing)  
- [License](#license)  

---

## 🌟 Features  
- **Responsive Design**: Fully responsive website using Bootstrap for seamless viewing on all devices.  
- **Dynamic Sections**: Includes Home, About Us, Services, and Plant Nursery sections.  
- **Modern Aesthetics**: Clean layout with engaging color schemes.  

---

## 🛠 Technologies Used  
- **HTML5**  
- **CSS3** (with a custom `style.css`)  
- **Bootstrap 4.5** for responsive grid and components  
- **Go** with **Fiber** for serving the website  

---

## ⚙️ Setup Instructions  

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/EthicalGopher/Lucky-Shrub.git
   cd Lucky-Shrub
   ```

2. **Install Dependencies**  
   Ensure you have [Go](https://golang.org/dl/) installed (version 1.22.9 or later).  
   Install Fiber by running:  
   ```bash
   go get -u github.com/gofiber/fiber/v2
   ```

3. **Run the Server**  
   Use the following code in the `main.go` file (located in the root directory):  
   ```go
   package main

   import (
       "github.com/gofiber/fiber/v2"
   )

   func main() {
       app := fiber.New()

       // Serve static files from the Frontend folder
       app.Static("/", "./Frontend")

       // Start the server
       app.Listen(":3000")
   }
   ```

   Run the server:  
   ```bash
   go run main.go
   ```

4. **Access the Website**  
   Open your browser and visit `http://localhost:3000` to view the website.  

---

## 🤝 Contributing  
Contributions are welcome! If you’d like to improve the project:  
1. Fork the repository.  
2. Create a new branch for your feature: `git checkout -b feature-name`.  
3. Commit your changes: `git commit -m "Add feature"`.  
4. Push to your branch: `git push origin feature-name`.  
5. Open a Pull Request.  

---

## 📜 License  
This project is licensed under the [MIT License](LICENSE).  

---

