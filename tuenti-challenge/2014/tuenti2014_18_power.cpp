
/*
 
Tuenti Challenge 4, Apr 2014, Challenge 18 - Power Password

Listing the directory we can find index.pyc, Python compiled file, when decompiled this is the most interesting parts:
...
 h, p = q.split(':')
...
 print 'Right!' if p.isdigit() and len(p) < 15 and pow(i[0], int(p), i[1]) == i[2] else 'Wrong!\n' + '\n'.join(map(str, i))
...
 
So have to solve the Discrete Logarithm Problem.
We'll use Baby-step Giant-step algorithm: http://en.wikipedia.org/wiki/Baby-step_giant-step
 
Most of this code is based in this implementation:
https://github.com/viralpoetry/Baby-step-giant-step
 

*/

#include <curlpp/cURLpp.hpp>
#include <curlpp/Easy.hpp>
#include <curlpp/Options.hpp>

#include <sstream>

#include <NTL/ZZ.h>

/* prime num = less collisions */
#define TABLE_SIZE  21677779

/* table computation & lookup*/
#define T 72428800
#define B 524288

int collisions = 0;

NTL_CLIENT

string url = "http://54.83.207.90:9083/index.py?";

int found = 0;
ZZ solution;

void raise_error()
{
    for(;;); // loop forever to avoid returning bad results
}

typedef struct hashlist
{
    long value;
    hashlist *p_coll;
    hashlist *p_next;
} HASHLIST;

typedef struct hashtable
{
    int size;
    // pointer to a pointer to a linked list
    HASHLIST **table;
} HASHTABLE;

HASHTABLE *htCreate( size_t size )
{
    HASHTABLE *fresh_table;

    if( (fresh_table = (HASHTABLE*)malloc(sizeof(HASHTABLE))) == NULL )
    {
	   return NULL;
    }

    if( (fresh_table->table = (HASHLIST**)malloc(sizeof(HASHLIST*)*size)) == NULL )
    {
	   return NULL;
    }

    int i;
    for( i = 0; i < size; i++ ) {
	   fresh_table->table[i] = NULL;
    }

    fresh_table->size = size;

    return fresh_table;
}

long htHash( HASHTABLE *t, ZZ value )
{
    unsigned long int result = 0;
    /*
    discrete logarithm results in pseudo random number
    so I can just apply the modulus table size
    */
    result = to_long(MulMod(value, 1, to_ZZ(t->size)));

    return result;
}

HASHLIST *htNewExponent( int value )
{
    HASHLIST *newList;

    if( (newList = (HASHLIST*)malloc(sizeof(HASHLIST))) == NULL )
    {
	   return NULL;
    }
    
    if( ( newList->value = value ) == NULL ) {
	   return NULL;
    }
    
    newList->p_coll = NULL;
    newList->p_next = NULL;
    
    return newList;
}

void htInsert( HASHTABLE *t, ZZ key, int value )
{
    int index = 0;
    HASHLIST *newExponent;
    HASHLIST *next ;
    HASHLIST *prev;

    index = htHash(t, key);
    next = t->table[index];
    
    /* collision */
    if( next != NULL )
    {
	   
	   /* walk to the end */
	   while( next->p_coll != NULL )
	   {
		  prev = next;
		  next = next->p_coll;
	   }
	   
	   newExponent = htNewExponent(value);
	   next->p_coll = newExponent;

	   collisions++;
    }
    else
    {
	   newExponent = htNewExponent(value);

	   /* start */
	   if( next == t->table[ index ] )
	   {
		  newExponent->p_next = next;
		  t->table[index] = newExponent;
        
	   }
	   /* end */
	   else if ( next == NULL )
	   {
		  prev->p_next = newExponent;	  
	   }
	   /* walking */
	   else
	   {
		  newExponent->p_next = next;
		  prev->p_next = newExponent;
	   }
    }
}

int findExponents( HASHTABLE *t, ZZ key, int x0, ZZ g, ZZ h, ZZ p )
{
    ZZ x;
    int index = 0;
    HASHLIST *pair;

    index = htHash(t, key);
    pair = t->table[index];

    /* shit happened */
    if( pair == NULL ) {
	   return NULL;
    }
    else
    {
	   x = (to_ZZ(pair->value) * B) + to_ZZ(x0);
	   
	   if( PowerMod(g, x, p) == h )
	   {
           found = 1;
           solution = to_ZZ(pair->value) * to_ZZ(B) + to_ZZ(x0);
		  return 1;
	   }

	   // walk collisions
	   while( pair->p_coll != NULL )
	   {
		  x = (to_ZZ(pair->p_coll->value) * B) + to_ZZ(x0);

		  if( PowerMod( g, x, p ) == h )
		  {
              found = 1;
              solution = to_ZZ(pair->p_coll->value) * to_ZZ(B) + to_ZZ(x0);
              return 1;
		  }
		  
		  pair = pair->p_coll;
	   }
	   
	   return pair->value;
    }
}

int main()
{
    string input;
    cin >> input;
        
    curlpp::Cleanup myCleanup;
    ostringstream read;
    read << curlpp::options::Url(url + input + ":0");
    string discard, base, exponent, residue;
    stringstream ss(read.str());
    getline(ss,discard,'\n');
    getline(ss,base,'\n');
    getline(ss,exponent,'\n');
    getline(ss,residue,'\n');
    
    HASHTABLE *tabulka = htCreate( TABLE_SIZE );

    if( tabulka == NULL )
    {
        raise_error();
    }
    
    ZZ p = to_ZZ(exponent.c_str());
    ZZ g = to_ZZ(base.c_str());
    ZZ h = to_ZZ(residue.c_str());
    g = g % p;
    
    int tbl;
    ZZ x;
    ZZ hgx1;
    ZZ gpowB;
    ZZ gx1  = h;
    ZZ gx0 = to_ZZ("1");

    /*
    precompute g^-B, there is no need 
    for computing it every iteration
    */
    PowerMod( gpowB, g, B, p );
    InvMod( gpowB, gpowB, p );
   
  
    int x1 = 1;
    for( ; x1 < T; x1++ )
    {
	   /* compute x1*g^-B */
	   MulMod( gx1, gx1, gpowB, p );

	   /* insert into table */
	   htInsert( tabulka, gx1, x1 );  
    }
  
    int x0 = 1;
    for(; x0 < B; x0++)
    {
	   /* (g^B)^x0 */
	   MulMod(gx0, gx0, g, p);

	   if(findExponents( tabulka, gx0, x0, g, h, p ) == 1)
		  break;
    }

    if (found) {
        cout << solution << endl;
    } else {
        raise_error();
    }
    
    return 0;
}
