import { AppProps } from "next/app";
import './global.css';
import Header from "../components/Header";
import Footer from "../components/Footer";

const RootLayout = ({children}) => {

    return (
      <>
      <html className="h-full bg-gray-100" lang="en">
        <body className="h-full">
          <div className="min-h-full">
           <Header />
            <main>
              <div className="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
              {children}
              </div>
            </main>
            <Footer />
          </div>
        </body>
      </html>
      </>
    );
}

export default RootLayout;
  