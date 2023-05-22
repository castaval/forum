'use client'

import Link from "next/link";


const CreatePage = () => {
    const handleSubmit = async (event: any) => {
        event.preventDefault();

        const data = {
            title: event.target.title.value,
            description: event.target.description.value,
            user_id: 1,
            channel_id: 1,
        };

        const JSONdata = JSON.stringify(data);

        const endpoint = 'http://localhost:4000/v1/threads';

        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSONdata,
        };

        const response = await fetch(endpoint, options);

        const result = await response.json();
        console.log(`Is this create thread response: ${result.data}`)
    };

    return (
        <>  
        <div className="py-8 px-4 mx-auto max-w-2xl lg:py-15">
            <h2 className="mb-4 text-xl font-bold text-gray-900 dark:text-white">Add a new thread</h2>
            <form onSubmit={handleSubmit}>
                <div className="grid gap-4 sm:grid-cols-2 sm:gap-6">
                    <div className="sm:col-span-2">
                        <label htmlFor="title" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Thread Name</label>
                        <input type="text" name="title" id="title" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Type thread name" required />
                    </div>
                    <div className="sm:col-span-2">
                        <label htmlFor="description" className="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Description</label>
                        <textarea id="description" rows={15} className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Your description here" required />
                    </div>
                </div>
                    <button type="submit" className="inline-flex items-center px-5 py-2.5 mt-4 sm:mt-6 text-sm font-medium text-center text-white bg-gray-700 rounded-lg focus:ring-4 focus:ring-primary-200 dark:focus:ring-gray-900 hover:bg-gray-800">
                        Add thread
                    </button>
            </form>
        </div>
        </>
    );
}

export default CreatePage;