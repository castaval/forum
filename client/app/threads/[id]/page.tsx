'use client';

import { useParams } from "next/navigation";

type Props = {
    params: {id: string};
    searchParams: { [key: string]: string | string[] | undefined}
}

const getThreadData = async (id: string) => {
    const thread = await fetch(`http://localhost:4000/v1/threads/${id}`)

    if (!thread.ok) {
        throw new Error('Failed to fetch thread data');
    }

    return thread.json();
}

const ThreadPage = async () => {
    const params = useParams();
    const threadData = await getThreadData(params.id)
    const thread = threadData.thread
    console.log(threadData)
    return (
        <>
        <div className="relative isolate overflow-hidden px-6 py-24 sm:py-32 lg:overflow-visible lg:px-0">
            <div className="mx-auto grid max-w-2xl grid-cols-1 gap-x-8 gap-y-16 lg:mx-0 lg:max-w-none lg:grid-cols-2 lg:items-start lg:gap-y-10">
                <div className="lg:col-span-2 lg:col-start-1 lg:row-start-1 lg:mx-auto lg:grid lg:w-full lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8">
                <div className="lg:pr-4">
                    <div className="lg:max-w-lg">
                    <h1 className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">{thread.title}</h1>
                    </div>
                </div>
                </div>
                <div className="lg:col-span-2 lg:col-start-1 lg:row-start-2 lg:mx-auto lg:grid lg:w-full lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8">
                    <div className="lg:pr-4">
                        <div className="max-w-xl text-base leading-7 text-gray-700 lg:max-w-lg">
                        <p>{thread.description}</p>
                        </div>
                    </div>
                </div>

                <div className="lg:col-span-2 lg:col-start-1 lg:row-start-3 lg:mx-auto lg:grid lg:w-full lg:max-w-7xl lg:grid-cols-2 lg:gap-x-8 lg:px-8">
                    <div className="col-span-full">
                        <label htmlFor="about" className="block text-sm font-medium leading-6 text-gray-900">Write</label>
                        <div className="mt-2">
                            <textarea id="about" name="about" rows={3} className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6" />
                        </div>
                    </div>

                    <div className="mt-6 flex items-center justify-start gap-x-6">
                        <button type="submit" className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Delete</button>
                    </div>
                    
                    <div className="mt-6 flex items-center justify-end gap-x-6">
                        <button type="button" className="text-sm font-semibold leading-6 text-gray-900">Cancel</button>
                        <button type="submit" className="rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600">Post</button>
                    </div>
                </div>
            </div>
        </div>
        </>
    );
}

export default ThreadPage;