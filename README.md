![EFES O](https://user-images.githubusercontent.com/27777864/236705863-ea61f2db-39ff-4c35-a4e4-f1f5e1d24b96.png)


## Query generator by template
### query generator starting from templates with simple ui


### Configuration 

In Efesto intallation folder <b>efesto_conf.json</b>

``` json
{
    "default_path":"[path to the working folder]"
}

```

#### Working folder 

    .
    ├── ...
    ├── [temporary file 'dbName']
    ├── templates            # working folder
    │   ├── [name of section]   # this name appare in the drop down
    │   │   ├── [dbName]        # the db name defined in the config file contains the template
    |   |   └── form            # the field appare in ui 
    │   └── config              # config file contains the list of db
    └── ...

#### config file

contains the list of db the name must be the same of the file in each section templates

``` json
{
    "databases":["postgres"]
}

```

#### template section

each template section contains the template files named like db and form file

FORM
``` json
[
    {
        "name":"key",
        "value":"text"
    },
    {
        "name":"it",
        "value":"text"
    },
    {
        "name":"en",
        "value":"text"
    }
]

```
the name rappresent the world in template using to replace this value.

``` sql
	----------- Query constant ---------------------

	INSERT INTO m_cons_constants (cons_id, cons_key, cons_name)
	SELECT nextval('seq_cons'),
		   '{{key}}',                          # key defined in form file
		   NULL WHERE NOT EXISTS
		(SELECT m_cons_constants.cons_id
		FROM m_cons_constants
		WHERE m_cons_constants.cons_key = '{{key}}' );
```    
