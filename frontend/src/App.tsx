import { useState } from 'react'
import './App.css'
import Store from './components/Store.tsx'
import CheckoutCart from './components/CheckoutCart.tsx';
import NavBar from './components/NavBar.tsx';
import { Product } from './types.ts';

function App() {
  const [isCheckoutOpen, setIsCheckoutOpen] = useState(false);
  const [cart, setCart] = useState<Product[]>(() => {
      const saved = localStorage.getItem("cart");
      return saved ? JSON.parse(saved) : [];
    });

  const onAddToCart = (product: Product) => {
    const updated = [...cart, product];
    setCart(updated);
    localStorage.setItem("cart", JSON.stringify(updated));
  };

  const onRemoveFromCart = (productId: number) => {
    const updated = cart.filter((p) => p.id !== productId);
    setCart(updated);
    localStorage.setItem("cart", JSON.stringify(updated));
  }

  const isInCart = (productId: number) => {
    return cart.some((p) => p.id === productId);
  }

  const onCartClick = () => {
    setIsCheckoutOpen((prev) => !prev);
  }

  const navItems = [
        { href: "/", display: "NYAB", text: "Not Your Average Benches"},
        { href: "/home", icon: "bx bx-home-alt", text: "Home" },
        { href: "/store", icon: "bx bx-grid-alt", text: "Store" },
        { href: "https://www.instagram.com/notyouraveragebenches/", icon: "bx bxl-instagram", text: "@notyouraveragebenches", external: true },
        { href: "/checkout", icon: "bx bx-cart", text: "View Cart", onClick: onCartClick},
        { href: "mailto:notyouraveragebenches@gmail.com", icon: "bx bx-envelope", text: <code>NotYourAverageBenches [at] gmail [dot] com</code>, external: true },
  ];

  return (
    <>
      <NavBar 
        navItems={navItems}
        onCartClick={onCartClick}
      />
      <Store 
        onAddToCart={onAddToCart}
        onRemoveFromCart={onRemoveFromCart}
        isInCart={isInCart}
      />
      <CheckoutCart
        isOpen={isCheckoutOpen}
        onClose={() => setIsCheckoutOpen(false)}
        cartItems={cart}
        onRemoveFromCart={onRemoveFromCart}
      />
    </>
  )
}

export default App
